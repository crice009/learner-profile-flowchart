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
    if (sequence_position > nodesData.length) sequence_position = 1; //right now the 'order' is 1-referenced, but that should change to 0-referenced in future versions
    console.log(sequence_position);
    var node = cy.filter('node[ order = "' + String(sequence_position) + '"]');
    console.log(node);
    cy.fit(node, 150); //150 is the relative pixel count around the selected item...
    document.getElementById('sequence').classList.add("on-page");
    document.getElementById('map').classList.remove("on-page");
};

function regress() {
    sequence_position -= 1;
    if (sequence_position < 1) sequence_position = 7; //right now the 'order' is 1-referenced, but that should change to 0-referenced in future versions
    console.log(sequence_position);
    var node = cy.filter('node[ order = "' + String(sequence_position) + '"]');
    console.log(node);
    cy.fit(node, 150); //150 is the relative pixel count around the selected item...
    document.getElementById('sequence').classList.add("on-page");
    document.getElementById('map').classList.remove("on-page");
};

function cy_fit() {
    cy.fit(100);
    sequence_position = 0;
    document.getElementById('sequence').classList.remove("on-page");
    document.getElementById('map').classList.add("on-page");
};