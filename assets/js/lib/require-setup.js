// file: js/require-setup.js
//
// Declare this variable before loading RequireJS JavaScript library
// To config RequireJS after it’s loaded, pass the below object into require.config();
var require = {
    baseUrl: '/js/lib',
    urlArgs: "v=0.0.1",
    paths: {
        'app': '../app',
        'jquery': 'jquery/jquery-2.1.3.min',
        'bootstrap': 'bootstrap/bootstrap.min',
        'jquery.treetable': 'jquery-treetable/jquery.treetable'
    },
    shim : {
        "bootstrap" : { "deps" :['jquery']},
    }
};
