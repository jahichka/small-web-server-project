window.onload = fetchSubjects();

function fetchSubjects() {
    fetch("http://127.0.0.1:8080/home?username=123&password=amina")
        .then(response => response.json())
        .then(response => {
            c
        });
}

