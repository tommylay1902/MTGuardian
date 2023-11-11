export function convertStringISO8601ToShortDate(dateString: string | null) {
  if (dateString == null) return;
  const date = new Date(dateString);
  const formattedDate = date.toLocaleDateString("en-US", {
    month: "2-digit",
    day: "2-digit",
    year: "numeric",
  });

  return formattedDate;
}
