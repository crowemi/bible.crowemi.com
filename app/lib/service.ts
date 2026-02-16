'use server';

import { Plan, PlanItem } from "./plan";
import { CONFIG } from "./config";
import { addToSet, removeFromSet } from "./firestore";
import { headers } from "next/headers";

async function getUserId(): Promise<string> {
    const headersList = await headers();
    const forwardedFor = headersList.get("x-forwarded-for");
    
    if (forwardedFor) {
        return forwardedFor.split(',')[0].trim();
    }
    
    return "127.0.0.1"; // Default/Fallback
}

async function GetPlan(id?: string) : Promise<Plan[] | Plan> {
    let plans: Response = null as any;
    if (id) {
        console.log("Fetching plan with ID: " + id);
        plans = await fetch(`${CONFIG.uri}/plan/${id}`)
    } else {
        console.log("Fetching all plans");
        plans = await fetch(`${CONFIG.uri}/plan`)
    }

    if (plans.ok) {
        const data = await plans.json() as Plan[]
        console.log(data)
        return data
    };
    return [];
}

async function GetPlanItem(planID: string) : Promise<PlanItem[]> {
    console.log("Plan ID:" + planID);
    const planItems = await fetch(`${CONFIG.uri}/plan/${planID}/item/`)
    if (planItems.ok) {
        const data = await planItems.json() as PlanItem[]
        console.log(data)
        return data
    } else {
        console.error("Failed to fetch plan items");
        console.error(planItems);
    };
    return [];
}

async function togglePlanItem(planItemId: string, isRead: boolean) {
    const userId = await getUserId();
    const collection = "plan_item";
    
    console.log(`Toggling plan item ${planItemId} for user ${userId} to ${isRead}`);
    
    try {
        if (isRead) {
            await addToSet(collection, planItemId, "CompletedBy", userId);
        } else {
            await removeFromSet(collection, planItemId, "CompletedBy", userId);
        }
    } catch (error) {
        console.error("Error toggling plan item:", error);
        throw error;
    }
}

export { GetPlan, GetPlanItem, togglePlanItem, getUserId };