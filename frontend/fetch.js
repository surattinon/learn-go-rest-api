// This code runs in the browser

async function fetchBooks() {
  try {
    const response = await fetch(`http://localhost:3000/books`);
    const data = await response.json();

    const usrSubmit = document.getElementById("usrSubmit");
    usrSubmit.addEventListener("submit", () => {
      if (usrSubmit === data.Title) {
        console.log("It's working");
      }
    });
    // data.forEach((book) => {
    //   console.log(
    //     book.find((first) => {
    //       return first.ID === 1;
    //     }),
    //   );
    // document.getElementById("bookTitle").innerHTML = book.Title;
    // document.getElementById("bookAuthor").innerHTML = book.Author;
    // });
  } catch (error) {
    console.error("Error fetching your books : ", error);
  }
}

fetchBooks();
