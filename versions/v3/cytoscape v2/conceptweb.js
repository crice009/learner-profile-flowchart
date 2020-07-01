var cy = cytoscape({
    container: document.getElementById('cy'),
    elements: [{
        data: { id: 'a' }
    }, {
        data: { id: 'b' }
    }, {
        data: { id: 'c' }
    }, {
        data: { id: 'd' }
    }, {
        data: { id: 'e' }
    }, {
        data: { id: 'ab', source: 'a', target: 'b' }
    }, {
        data: { id: 'ac', source: 'a', target: 'c' }
    }, {
        data: { id: 'ae', source: 'a', target: 'e' }
    }, {
        data: { id: 'bc', source: 'b', target: 'c' }
    }, {
        data: { id: 'bd', source: 'b', target: 'd' }
    }, {
        data: { id: 'ce', source: 'c', target: 'e' }
    }],
    style: [{ //http://manual.cytoscape.org/en/stable/Styles.html#introduction-to-style
        selector: 'node',
        style: {
            shape: 'hexagon',
            'background-color': 'green',
            label: 'data(id)'
        }
    }],
    layout: {
        name: 'circle'
    }

});