function Verification() {
    var user = document.getElementById('username').value;
    var pass = document.getElementById('password').value;
    console.log("Username:", user);
    console.log("Password:", pass);

    fetch(`http://127.0.0.1:8080/dashboard?username=${user}&password=${pass}`)
        .then(response => {
            if (response.ok) {
                console.log("Response status:", response.status);
                return response.json();
            } else {
                console.error("Response error:", response.status);
                throw new Error('Network response was not ok.');
            }
        })
        .then(data => {
            console.log("Data received:", data);
            const main=document.getElementById("main");
            main.innerHTML=" "
            for (const unit of data){
                main.innerHTML+=`
                <div class="data">
                ${unit.plastenik}, ${unit.biljka}
                </div>
               `
            }
            const ButtonsSearch=document.getElementById("submitSearch")
            ButtonsSearch.innerHTML=""
            ButtonsSearch.innerHTML+=`
            <section>
            <button id='add' onclick="Add()">Dodaj</button>
            </section>
            <section>
            <input type="text" id="searchInput" placeholder="Search">
            </section>
            <section>
            <button id='search' onclick="Search()">Trazi</button>
            </section>
            <div id='results'>
            </section>
            `
        })
        .catch(error => {
            console.error('Fetch error:', error);
            alert('An error occurred while fetching data.');
        });

    return false;
}

async function Add() {
    var ButtonsSearch= document.getElementById("submitSearch")
    ButtonsSearch.innerHTML += `
    <section id='forma'>
        <input type="text" id="reqPlastenik" placeholder="Naziv mjesta"> <br>
        <input type="text" id="reqBiljka" placeholder="Biljka"> <br>
        <div>
            <button onclick="newEntry()">DODAJ</button>
            <button onclick="overlayOff()">PONISTI</button>
        </div>
    </section>
    `
}
function overlayOff(){
    var div=document.getElementById("forma")
    div.innerHTML=""
}

function newEntry(){
    var username='123'
    var password='amina'
    var plastenik=document.getElementById("reqPlastenik").value;
    var biljka=document.getElementById("reqBiljka").value;

    fetch('http://127.0.0.1:8080/add', {
        method: 'POST',
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ "Username": username, "Plastenik": plastenik, "Biljka": biljka})
    })
        .then(response => {
            console.log(response)
            if(response.ok){
            var main = document.getElementById("main")
            main.innerHTML += `<div>
            ${plastenik}, ${biljka}
            </div>
            `;
            overlayOff();
        }
        });
} 

function Search() {
    var query = document.getElementById("searchInput").value.toLowerCase();
    console.log(query)
    var elements = document.getElementsByClassName("data");
    console.log(elements)
    var results = document.getElementById("results")
    results.style=`

    `
    results.innerHTML=""
    for (var i = 0; i < elements.length; i++) {
        var text = elements[i].textContent.toLowerCase();
        console.log(text);
        if (text.includes(query)) {
            console.log(text);
            results.innerHTML+=`
            <div>
            ${text}
            </div>` 
        }
    }
}

