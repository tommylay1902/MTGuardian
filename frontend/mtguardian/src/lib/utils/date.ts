import { DateTime } from "luxon";

export function convertStringISO8601ToShortDate(dateString: string | null) {
  if (dateString == null) return;
  const parsedDate = DateTime.fromISO(dateString, { zone: "utc" }).toFormat(
    "MM-dd-yyyy"
  );
  return parsedDate;
}
