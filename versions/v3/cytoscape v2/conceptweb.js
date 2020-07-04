var cy = cytoscape({
    container: document.getElementById('cy'),
    elements: {
        nodes: [
            {data: { id: 'n1', name:'Fab Lab Skills' }}, 
            {data: { id: 'n2', name:'3D Design' }}, 
            {data: { id: 'n3', name:'2D Design' }}, 
            {data: { id: 'n4', name:'Fab Lab Mindset' }},
            {data: { id: 'n5', name:'Computational Thinking' }}, 
            {data: { id: 'n6', name:'Electronics' }}, 
            {data: { id: 'n7', name:'Microcontrollers' }}
        ],
        edges: [

            {data: { id: '1-2', source: 'n1', target: 'n2' }}, 
            {data: { id: '1-3', source: 'n1', target: 'n3' }}, 
            {data: { id: '1-4', source: 'n1', target: 'n4' }}, 
            {data: { id: '1-5', source: 'n1', target: 'n5' }}, 
            {data: { id: '1-6', source: 'n1', target: 'n6' }}, 
            {data: { id: '1-7', source: 'n1', target: 'n7' }}
        ]},
    style: [{ //http://manual.cytoscape.org/en/stable/Styles.html#introduction-to-style
        selector: 'node',
        style: {
            shape: 'concentric',
            'background-color': 'green',
            'color': 'white',
            label: 'data(name)'
        }
    }],
    layout: {
        name: 'concentric', //https://js.cytoscape.org/#layouts/concentric
        spacingFactor: 5,
        concentric: function( node ){
          return node.degree();
        },
        levelWidth: function( nodes ){
          return 1;
        }
      },


});