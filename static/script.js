// function for message box when deleting
function confirmDelete(id, judul) {
    Swal.fire({
        title: `Are you sure delete ${judul}?`, 
        text: 'You will not be able to recover this item!',
        icon: 'warning',
        showCancelButton: true,
        confirmButtonColor: '#d33',
        cancelButtonColor: '#3085d6',
        confirmButtonText: 'Yes, delete it!'
    }).then((result) => {
        if (result.isConfirmed) {
            document.getElementById('deleteForm' + id).submit(); 
        }
    });
    return false;
}

// function for rendering star in table
function renderStars(rating) {
    let stars = '';
    for (let i = 0; i < 5; i++) {
        if (i < rating) {
            stars += '<i class="fas fa-star gold-star"></i>'; 
        } else {
            stars += '<i class="far fa-star gold-star"></i>'; 
        }
    }
    return stars;
}

document.querySelectorAll('.stars').forEach(function(starElement) {
    const rating = parseInt(starElement.textContent); 
    starElement.innerHTML = renderStars(rating); 
});

$(document).ready(function() {
    $('#dataTable').DataTable({
        "searching": true,
    });
});