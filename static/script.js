function confirmDelete(id) {
    // Display confirmation dialog using SweetAlert
    Swal.fire({
        title: 'Are you sure?',
        text: 'You will not be able to recover this item!',
        icon: 'warning',
        showCancelButton: true,
        confirmButtonColor: '#d33',
        cancelButtonColor: '#3085d6',
        confirmButtonText: 'Yes, delete it!'
    }).then((result) => {
        if (result.isConfirmed) {
            // If user confirms deletion, submit the form
            document.getElementById('deleteForm').submit();
        }
    });
    // Prevent default form submission
    return false;
}
