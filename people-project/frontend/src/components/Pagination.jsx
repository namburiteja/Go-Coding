function Pagination({
  page,
  totalPages,
  onPrevious,
  onNext,
}) {
  return (
    <div className="pagination">

      <div className="pagination-info">
        Showing page {page} of {totalPages}
      </div>

      <div className="pagination-buttons">
        <button
          onClick={onPrevious}
          disabled={page === 1}
        >
          ← Previous
        </button>

        <span className="page-number">
          {page}
        </span>

        <button
          onClick={onNext}
          disabled={page === totalPages}
        >
          Next →
        </button>
      </div>

    </div>
  );
}

export default Pagination;