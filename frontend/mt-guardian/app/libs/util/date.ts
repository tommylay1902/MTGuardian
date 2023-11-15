import { DateTime } from "luxon";

export default function convertDate(date: string) {
  //   const convertedDate = moment(date).utc().format("YYYY-MM-DD");
  //   return convertedDate;
  const parsedDate = DateTime.fromISO(date, { zone: "utc" }).toFormat(
    "yyyy-MM-dd"
  );
  return parsedDate;
}
