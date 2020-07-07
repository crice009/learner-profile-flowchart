var cy = cytoscape({
    container: document.getElementById('cy'),
    elements: {
        nodes: [
            { data: { id: 'n1', order: 1, name: 'Fab Lab Skills', href: 'https://prezi.com/view/zBqbXsXRkA4yfSB1OiAT/' } },
            { data: { id: 'n2', order: 4, name: '3D Design' } },
            { data: { id: 'n3', order: 3, name: '2D Design' } },
            { data: { id: 'n4', order: 2, name: 'Fab Lab Mindset' } },
            { data: { id: 'n5', order: 5, name: 'Computational Thinking' } },
            { data: { id: 'n6', order: 6, name: 'Electronics' } },
            { data: { id: 'n7', order: 7, name: 'Microcontrollers' } }
        ],
        edges: [

            { data: { id: '1-2', source: 'n1', target: 'n2' } },
            { data: { id: '1-3', source: 'n1', target: 'n3' } },
            { data: { id: '1-4', source: 'n1', target: 'n4' } },
            { data: { id: '1-5', source: 'n1', target: 'n5' } },
            { data: { id: '1-6', source: 'n1', target: 'n6' } },
            { data: { id: '1-7', source: 'n1', target: 'n7' } }
        ]
    },
    style: [{ //http://manual.cytoscape.org/en/stable/Styles.html#introduction-to-style
        selector: 'node',
        style: {
            shape: 'concentric',
            'background-color': 'rgb(94,130,23)',
            'font-family': '"Oswald", sans-serif',
            'color': 'white',
            label: 'data(name)'
        }
    }],
    layout: {
        name: 'concentric', //https://js.cytoscape.org/#layouts/concentric
        spacingFactor: 5, // makes it so the nodes aren't on top of each other (multiplies edge lengths)
        concentric: function(node) {
            return node.degree();
        },
        levelWidth: function(nodes) {
            return 1;
        },
    }
});

cy.fit(100); //make it look nice when it first loads - right now this means the whole map is on screen with 'padding'