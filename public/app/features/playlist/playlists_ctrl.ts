import _ from 'lodash';
import coreModule from '../../core/core_module';

export class PlaylistsCtrl {
  playlists: any;
  navModel: any;

  /** @ngInject */
  constructor(private $scope, private backendSrv, navModelSrv) {
    this.navModel = navModelSrv.getNav('dashboards', 'playlists', 0);

    backendSrv.get('/api/playlists').then(result => {
      this.playlists = result.map(item => {
        item.startUrl = `playlists/play/${item.id}`;
        return item;
      });
    });
  }

  removePlaylistConfirmed(playlist) {
    _.remove(this.playlists, { id: playlist.id });

    this.backendSrv.delete('/api/playlists/' + playlist.id).then(
      () => {
        this.$scope.appEvent('alert-success', ['删除播放列表', '']);
      },
      () => {
        this.$scope.appEvent('alert-error', ['无法删除播放列表', '']);
        this.playlists.push(playlist);
      }
    );
  }

  removePlaylist(playlist) {
    this.$scope.appEvent('confirm-modal', {
      title: '删除',
      text: '确定需要删除此' + playlist.name + '面板吗?',
      yesText: '删除',
      icon: 'fa-trash',
      onConfirm: () => {
        this.removePlaylistConfirmed(playlist);
      },
    });
  }
}

coreModule.controller('PlaylistsCtrl', PlaylistsCtrl);
