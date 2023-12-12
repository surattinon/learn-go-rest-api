async function fetchBooks() {
  try {
    const response = await fetch(`http://localhost:3000/`);

    const data = await response.json();
    for (const book of data) {
      console.log(book);
    }
    document.getElementById("myBook").innerHTML = book.title;
  } catch (error) {
    console.error("Error fetching your books : ", error);
    return null;
  }
}

fetchBooks();
