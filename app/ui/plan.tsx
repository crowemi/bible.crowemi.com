import Link from "next/link";
import { Plan } from "../lib/plan";

export default function PlanComponent(plan: Plan) {
  const formatDate = (s?: string) => {
    if (!s) return "—";
    try {
      return new Date(s).toLocaleDateString(undefined, { month: "short", day: "numeric", year: "numeric" });
    } catch (_) {
      return s;
    }
  };

  const start = formatDate(plan.PeriodStart);
  const end = formatDate(plan.PeriodEnd);

  return (
    <div className="my-8 mx-auto max-w-4xl px-6">
      <ul role="list" className="mt-8 w-full">
        <li
          key={plan.PlanID}
          className="w-full bg-white dark:bg-gray-900 border border-gray-100 dark:border-gray-800 rounded-lg p-6 shadow-sm"
        >
          <div className="flex flex-col sm:flex-row sm:items-center sm:justify-between gap-4">
            <div className="flex items-start gap-4">
              <div className="flex-shrink-0 rounded-full bg-indigo-50 dark:bg-indigo-900 p-3">
                <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" className="size-6">
                  <path strokeLinecap="round" strokeLinejoin="round" d="M12 6.042A8.967 8.967 0 0 0 6 3.75c-1.052 0-2.062.18-3 .512v14.25A8.987 8.987 0 0 1 6 18c2.305 0 4.408.867 6 2.292m0-14.25a8.966 8.966 0 0 1 6-2.292c1.052 0 2.062.18 3 .512v14.25A8.987 8.987 0 0 0 18 18a8.967 8.967 0 0 0-6 2.292m0-14.25v14.25" />
                </svg>
              </div>

              <div>
                <h3 className="text-lg font-semibold text-gray-900 dark:text-white">{plan.Name}</h3>
                <p className="mt-1 text-sm text-gray-500 dark:text-gray-400">{plan.Description}</p>

                <div className="mt-3 flex flex-wrap gap-2">
                  <span className="inline-flex items-center rounded-md bg-indigo-100 px-2 py-1 text-xs font-medium text-indigo-600 dark:bg-indigo-400/10 dark:text-indigo-400">Duration: {plan.Duration}</span>
                  <span className="inline-flex items-center rounded-md bg-gray-100 px-2 py-1 text-xs font-medium text-gray-600 dark:bg-white/5 dark:text-gray-300">Starts: {start}</span>
                  <span className="inline-flex items-center rounded-md bg-gray-100 px-2 py-1 text-xs font-medium text-gray-600 dark:bg-white/5 dark:text-gray-300">Ends: {end}</span>
                </div>
              </div>
            </div>

            <div className="ml-0 sm:ml-4 flex items-center gap-3">
              <Link
                href={`/plan/${plan.PlanID}`}
                className="inline-flex items-center rounded-md bg-indigo-600 px-4 py-2 text-sm font-medium text-white hover:bg-indigo-700"
              >
                Open
              </Link>
            </div>
          </div>
        </li>
      </ul>
    </div>
  );
}
