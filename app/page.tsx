import { Plan } from "./lib/plan";
import PlanComponent from "./ui/plan";
import {GetPlan} from "./lib/service";

function renderPlanComponent(plans: Plan[]) {
  return plans.map((plan) => (
    <PlanComponent key={plan.PlanID} {...plan} />
  ));
}

export default async function Home() {
  const plans = await GetPlan() as Plan[];

  return (
    <div className="mx-auto dark:bg-gray-900">
      <div>
        <div className="bg-white px-6 py-12 sm:py-24 lg:px-8 dark:bg-gray-900">
          <div className="mx-auto max-w-2xl text-center">
            <h2 className="text-5xl font-semibold tracking-tight text-gray-900 sm:text-7xl dark:text-white">
              Sola Scriptura
            </h2>
            <p className="mt-8 text-lg font-medium text-pretty text-gray-500 sm:text-xl/8 dark:text-gray-400">
              The{" "}
              <span className="emphasis font-medium text-indigo-400">
                blessed
              </span>{" "}
              man delights in the law of the Lord, and on his law he{" "}
              <span className="emphasis font-medium text-indigo-400">
                meditates
              </span>{" "}
              day and night. (Psalm 1:2)
            </p>
          </div>
        </div>
        { renderPlanComponent(plans) }
      </div>
    </div>
  );
}
