function SortControls({
  sortBy,
  sortOrder,
  onSortByChange,
  onSortOrderChange,
}) {
  const handleSortByChange = (event) => {
    const value = event.target.value;

    onSortByChange(value);

    if (value === "nameFirst") {
      onSortOrderChange("asc");
    } else {
      onSortOrderChange("asc");
    }
  };

  const handleSortOrderChange = (event) => {
    onSortOrderChange(event.target.value);
  };

  return (
    <div className="sort-controls">

      <label htmlFor="sortBy">
        Sort By
      </label>

      <select
        id="sortBy"
        value={sortBy}
        onChange={handleSortByChange}
      >
        <option value="">
          Default
        </option>

        <option value="nameFirst">
          First Name
        </option>

        <option value="birthYear">
          Birth Year
        </option>

        <option value="height">
          Height
        </option>
      </select>

      <select
        value={sortOrder}
        onChange={handleSortOrderChange}
        disabled={!sortBy}
      >
        <option value="asc">
          {sortBy === "nameFirst"
            ? "A → Z"
            : "Min → Max"}
        </option>

        <option value="desc">
          {sortBy === "nameFirst"
            ? "Z → A"
            : "Max → Min"}
        </option>
      </select>

    </div>
  );
}

export default SortControls;