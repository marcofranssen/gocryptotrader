'use strict';

angular.module('myApp.home', ['ngRoute'])

.config(['$routeProvider', function($routeProvider) {
    $routeProvider.when('/', {
        templateUrl: '/views/home/home.html',
        controller: 'HomeController'
    });
}])

.controller('HomeController', function($scope, $rootScope, $timeout, webSocket) {
    $scope.loaded = false;

    $rootScope.$on('CurrencyChanged', function(event, args) {
        $scope.currency = args.Currency;
        $scope.exchange = args.Exchange;
        $scope.loaded = true;
    });

    $timeout(function() {
        if ($scope.currency) {
            $scope.loaded = true;
        } else {
            $scope.loadFailed = true;
        }
    }, 10000);
});