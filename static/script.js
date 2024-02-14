function confirmDelete(id, judul) {
    // Display confirmation dialog using SweetAlert
    Swal.fire({
        title: `Are you sure delete ${judul}?`, // Use the book title from the data attribute
        text: 'You will not be able to recover this item!',
        icon: 'warning',
        showCancelButton: true,
        confirmButtonColor: '#d33',
        cancelButtonColor: '#3085d6',
        confirmButtonText: 'Yes, delete it!'
    }).then((result) => {
        if (result.isConfirmed) {
            // If user confirms deletion, submit the form
            document.getElementById('deleteForm' + id).submit(); // Use unique form ID
        }
    });
    // Prevent default form submission
    return false;
}

function renderStars(rating) {
    let stars = '';
    for (let i = 0; i < 5; i++) {
        if (i < rating) {
            stars += '<i class="fas fa-star gold-star"></i>'; // Ikuti kelas Font Awesome untuk ikon bintang yang diisi
        } else {
            stars += '<i class="far fa-star gold-star"></i>'; // Ikuti kelas Font Awesome untuk ikon bintang kosong
        }
    }
    return stars;
}

// Ambil semua elemen dengan kelas "stars" dan ubah isi HTML-nya menjadi ikon bintang
document.querySelectorAll('.stars').forEach(function(starElement) {
    const rating = parseInt(starElement.textContent); // Ambil nilai rating dari konten elemen
    starElement.innerHTML = renderStars(rating); // Ubah isi HTML elemen menjadi ikon bintang
});

$(document).ready(function() {
    $('#dataTable').DataTable({
        "searching": true,
        // Atur opsi lainnya sesuai kebutuhan Anda
    });
});