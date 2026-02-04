import {GetPlanItem, GetPlan} from "../../lib/service";
import PlanComponent from "../../ui/plan";
import PlanItemComponent from "../../ui/planItem";
import Link from "next/link";
import { Plan, PlanItem } from "../../lib/plan";


function renderPlanItem(planItems: PlanItem[]) {
  if (planItems.length === 0) {
    return (
      <div className="flex items-center justify-center py-16">
        <div className="text-center bg-white dark:bg-gray-900 border border-gray-100 dark:border-gray-800 rounded-lg px-6 py-10 shadow-sm max-w-xl mx-auto">
          <div className="mx-auto flex h-12 w-12 items-center justify-center rounded-full bg-indigo-50 dark:bg-indigo-900">
            <svg xmlns="http://www.w3.org/2000/svg" className="h-6 w-6 text-indigo-600 dark:text-indigo-300" fill="none" viewBox="0 0 24 24" stroke="currentColor" aria-hidden="true">
              <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={1.5} d="M12 6v6l4 2" />
            </svg>
          </div>
          <h3 className="mt-4 text-2xl font-semibold text-gray-900 dark:text-white">Take it easy</h3>
          <p className="mt-2 text-sm text-gray-500 dark:text-gray-400">No reading items today. Check back later or choose a different plan.</p>
          <div className="mt-6">
            <Link href="/" className="my-4 rounded-md bg-indigo-50 px-2.5 py-1.5 text-sm font-semibold text-indigo-600 shadow-xs hover:bg-indigo-100 dark:bg-indigo-500/20 dark:text-indigo-400 dark:shadow-none dark:hover:bg-indigo-500/30">Home</Link>
          </div>
        </div>
      </div>
    );
  }
  
  return planItems.map((planItem) => (
    <PlanItemComponent key={planItem.PlanItemID} {...planItem} />
  ));
}

export default async function Page({
  params,
}: {
  params: Promise<{ id: string }>
}) {

  const id = (await params).id
  const plan = await GetPlan(id) as Plan;
  const planItems = await GetPlanItem(id);

  return (
    <div className="mx-auto dark:bg-gray-900">
      <div>
        <div className="bg-white px-6 py-12 sm:py-24 lg:px-8 dark:bg-gray-900">
          <div className="mx-auto max-w-2xl text-center">
            <h2 className="text-5xl font-semibold tracking-tight text-gray-900 sm:text-7xl dark:text-white">
              {plan.Name}
            </h2>
            <p className="mt-8 text-lg font-medium text-pretty text-gray-500 sm:text-xl/8 dark:text-gray-400">
              {plan.Description}
            </p>
          </div>
        </div>
        {renderPlanItem(planItems)}
      </div>
    </div>
  );
}
