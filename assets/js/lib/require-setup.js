// file: js/require-setup.js
//
// Declare this variable before loading RequireJS JavaScript library
// To config RequireJS after it’s loaded, pass the below object into require.config();
var require = {
    baseUrl: 'js/lib',
    urlArgs: "v=0.0.1",
    paths: {
        'app': '../app',
        'jquery': 'jquery/jquery-1.11.2',
        'bootstrap': 'bootstrap/bootstrap',
        'jquery.treetable': 'jquery-treetable/jquery.treetable'
    },
    shim : {
        "bootstrap" : { "deps" :['jquery']},
    }
};
