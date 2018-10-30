gameTable = document.querySelector("tbody")

/*
Use the browsers `fetch` API to make a GET call to /bird
We expect the response to be a JSON list of birds, of the
form :
[
    {"species":"...","description":"..."},
    {"species":"...","description":"..."}
]
*/
fetch("/games")
    .then(response => response.json())
    .then(gameList => {
    //Once we fetch the list, we iterate over it
    gameList.forEach(game => {
        // Create the table row
        row = document.createElement("tr")

        alert(game.id);
        console.log("ITERAT");
        // Create the table data elements for the species and
                // description columns
        id = document.createElement("td")
        id.innerHTML = game.id
        home = document.createElement("td")
        home.innerHTML = game.home
        away = document.createElement("td")
        away.innerHTML = game.away
        ts = document.createElement("td")
        ts.innerHTML = game.ts

        // Add the data elements to the row
        row.appendChild(id)
        row.appendChild(home)
        row.appendChild(away)
        row.appendChild(ts)
        // Finally, add the row element to the table itself
        gameTable.appendChild(row)
    })
})