'use client';

import { useState } from "react";
import { PlanItem } from "../lib/plan";

function PlanItemComponent(props: PlanItem) {
  const [isRead, setIsRead] = useState(false);


  function markRead() {
    setIsRead((previousState) => { return !previousState; })
  }

  function renderReadStatus() {
    if (isRead) {
      return (
        <span className="inline-flex items-center rounded-md bg-green-100 px-2 py-1 my-4 mx-1 text-xs font-medium text-green-600 dark:bg-green-400/10 dark:text-green-400">
          Read
        </span>
      );
    }
    return (
      <span className="inline-flex items-center rounded-md bg-indigo-100 px-2 py-1 my-4 mx-1 text-xs font-medium text-indigo-600 dark:bg-indigo-400/10 dark:text-indigo-400">
        Unread
      </span>
    );
  }
  function renderReadButtonType() {
    if (isRead) {
      return "Mark Unread";
    }
    return "Mark Read";
  }

  return (
    <div>
      <ul role="list" className="divide-y divide-gray-100 dark:divide-white/5">
        <li className="mx-auto py-6 px-4 max-w-4xl">
          <div className="flex flex-col sm:flex-row sm:items-center sm:justify-between gap-4">
            <div className="flex items-start gap-4">
              <div className="flex-shrink-0 rounded-full bg-indigo-50 dark:bg-indigo-900 p-3">
              <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" className="size-6">
                <path strokeLinecap="round" strokeLinejoin="round" d="M17.593 3.322c1.1.128 1.907 1.077 1.907 2.185V21L12 17.25 4.5 21V5.507c0-1.108.806-2.057 1.907-2.185a48.507 48.507 0 0 1 11.186 0Z" />
              </svg>

              </div>

              <div>
                <a href={`plan/xxxx/item/${props.PlanItemID}`} className="group inline-block">
                  <h3 className="text-2xl font-semibold tracking-tight text-gray-900 dark:text-white group-hover:text-indigo-600">
                    {props.Book.Name} <span className="ml-2 text-sm font-medium text-gray-500 dark:text-gray-400">Chapter {props.Chapter.Number}</span>
                  </h3>
                </a>

                <p className="mt-1 text-sm text-gray-500 dark:text-gray-400">{props.Summary ?? props.Chapter?.Summary ?? props.Book?.Summary ?? ''}</p>

                <div className="mt-3 flex flex-wrap items-center gap-2">
                  <span className="inline-flex items-center rounded-md bg-gray-100 px-2 py-1 text-xs font-medium text-gray-600 dark:bg-white/5 dark:text-gray-300">{props.ReadDate}</span>
                  {renderReadStatus()}
                </div>
              </div>
            </div>

            <div className="ml-0 sm:ml-4 flex items-center gap-3">
              <button
                onClick={markRead}
                type="button"
                className="my-0 rounded-md bg-indigo-50 px-2.5 py-1.5 text-sm font-semibold text-indigo-600 shadow-xs hover:bg-indigo-100 dark:bg-indigo-500/20 dark:text-indigo-400 dark:shadow-none dark:hover:bg-indigo-500/30"
              >
                {renderReadButtonType()}
              </button>
            </div>
          </div>

          <p className="mt-4 text-base text-gray-700 dark:text-gray-300">
            Lorem Ipsum is simply dummy text of the printing and typesetting
            industry. Lorem Ipsum has been the industry's standard dummy text
            ever since the 1500s, when an unknown printer took a galley of type
            and scrambled it to make a type specimen book. It has survived not
            only five centuries, but also the leap into electronic typesetting,
            remaining essentially unchanged. It was popularised in the 1960s
            with the release of Letraset sheets containing Lorem Ipsum passages,
            and more recently with desktop publishing software like Aldus
            PageMaker including versions of Lorem Ipsum.
          </p>
        </li>
      </ul>
    </div>
  );
}

export default PlanItemComponent;
