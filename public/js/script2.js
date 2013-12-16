angular.module('script', [], function($routeProvider, $locationProvider) {
     $locationProvider.html5Mode(true);
});

function Ctrl($scope, $location) {
    $scope.language = 'pt-BR';
    if (window.location.pathname == '/en' || $location.path() == '/en') {
        $scope.language = 'en-US';
    }

    $scope.changeLanguage = function(lang) {
        $scope.language = lang;
        if (lang == 'en-US') {
            $location.path('/en');
        } else {
            $location.path('/');
        }
    };
}
