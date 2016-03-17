var UserController = function($scope, $state, $uibModal, formUtils, locationUtils, userService, authService, userModel) {

    var updateUser = function(user, options) {
        userService.updateCurrentUser(user)
            .then(function() {
                if (options.signout) {
                    authService.logout();
                }
            });
    };

    $scope.userOriginal = angular.copy(userModel.user);

    $scope.user = userModel.user;

    $scope.confirmUpdate = function(user, usernameField) {
        if (usernameField.$dirty) {
            var params = {
                title: 'Reauthentication Required',
                message: 'Changing your username to ' + user.username + ' will require you to reauthenticate. Is that OK?'
            };
            var modalInstance = $uibModal.open({
                templateUrl: 'common/modules/dialog/confirm/dialog.confirm.tpl.html',
                controller: 'DialogConfirmController',
                size: 'sm',
                resolve: {
                    params: function () {
                        return params;
                    }
                }
            });
            modalInstance.result.then(function() {
                updateUser(user, { signout : true });
            }, function () {
                // do nothing
            });
        } else {
            updateUser(user, { signout : false });
        }
    };

    $scope.navigateToPath = locationUtils.navigateToPath;

    $scope.hasError = formUtils.hasError;

    $scope.hasPropertyError = formUtils.hasPropertyError;

};

UserController.$inject = ['$scope', '$state', '$uibModal', 'formUtils', 'locationUtils', 'userService', 'authService', 'userModel'];
module.exports = UserController;