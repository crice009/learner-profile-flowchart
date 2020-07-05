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
};

function regress() {

};