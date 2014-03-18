angular.module('app', ['angularFileUpload'])
    .controller('MusicUploadController', function($scope, $fileUploader) {
        'user strict';

        var uploader = $scope.uploader = $fileUploader.create({
            scope: $scope,
            url: '/music/upload'
        });

        uploader.filters.push(function(item) {
            if (item.type == 'audio/mp3') {
                return true;
            } else {
                return false;
            }
        });

        uploader.filters.push(function(item) {
            if (item.size > 100000000) {
                return false;
            } else {
                return true;
            }
        });
    });
