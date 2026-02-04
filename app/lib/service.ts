'use server';

import { Plan, PlanItem } from "./plan";
import { CONFIG } from "./config";


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

export { GetPlan, GetPlanItem };