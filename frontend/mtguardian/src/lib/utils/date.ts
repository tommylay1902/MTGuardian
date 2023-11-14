import { DateTime } from "luxon";

export function convertStringISO8601ToShortDate(dateString: string | null) {
  if (dateString == null) return;
  const parsedDate = DateTime.fromISO(dateString, { zone: "utc" }).toFormat(
    "MM-dd-yyyy"
  );

  return parsedDate;
}

export function convertStringISO8601ToDateHtmlInput(date: string) {
  const parsedDate = DateTime.fromISO(date, { zone: "utc" }).toFormat(
    "yyyy-MM-dd"
  );
  return parsedDate;
}

export function convertDateHtmlInputStringToISO8601(dateString: string) {
  const parsedDate = DateTime.fromFormat(dateString, "yyyy-MM-dd");

  return parsedDate.toFormat("yyyy-MM-dd'T'HH:mm:ss.SSS'Z'");
}
