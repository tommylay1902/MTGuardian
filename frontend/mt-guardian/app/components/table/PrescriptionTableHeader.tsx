import React from "react";
type Props = {
  tableHeaders: string[];
  tableHeaderExclusions: string[];
};
const PrescriptionTableHeader: React.FC<Props> = ({
  tableHeaders,
  tableHeaderExclusions,
}) => {
  return (
    <thead className="w-full text-xl text-gray-700 uppercase bg-gray-50 dark:bg-gray-700 dark:text-gray-400">
      <tr>
        {tableHeaders.map((h) => {
          if (!tableHeaderExclusions.includes(h)) {
            return (
              <th scope="col" className="px-4 py-3" key={h}>
                {" "}
                {/* Decreased the padding */}
                {h}
              </th>
            );
          }
        })}
        <th scope="col" className="px-6 py-3">
          Edit/Delete
        </th>{" "}
        {/* Keep the last column with the original padding */}
      </tr>
    </thead>
  );
};

export default PrescriptionTableHeader;
