cy.on('tap', 'node', function() {
    try { // your browser may block popups
        window.open(this.data('href'));
    } catch (e) { // fall back on url change
        window.location.href = this.data('href');
    }
});

var sequence_position = 0;

function advance() {
    sequence_position += 1;
    if (sequence_position > 7) sequence_position = 0;

    var node = cy.filter('node[order = ' + String(sequence_position) + ']');
    console.log(node);
    cy.fit(node, 150); //150 is the relative pixel count around the selected item...
    document.getElementById('sequence').classList.add("on-page");
    document.getElementById('map').classList.remove("on-page");
};

function regress() {
    sequence_position -= 1;
    if (sequence_position < 0) sequence_position = 7;

    var node = cy.filter('node[order = ' + String(sequence_position) + ']');
    console.log(node);
    cy.fit(node, 150); //150 is the relative pixel count around the selected item...
    document.getElementById('sequence').classList.add("on-page");
    document.getElementById('map').classList.remove("on-page");
};

function cy_fit() {
    cy.fit();
    document.getElementById('sequence').classList.remove("on-page");
    document.getElementById('map').classList.add("on-page");
};