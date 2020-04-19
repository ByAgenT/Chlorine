/* eslint-disable @typescript-eslint/camelcase */
class SpotifyPlayer {
  private readonly player: Spotify.SpotifyPlayer;

  constructor(player) {
    this.player = player;
  }

  async getCurrentState(): Promise<Spotify.PlaybackState | null> {
    if (this.player !== undefined) {
      return this.player.getCurrentState();
    }
  }

  async onPlayerReady(callback: Spotify.PlaybackInstanceListener): Promise<void> {
    if (this.player !== undefined) {
      this.player.removeListener('ready');
      this.player.addListener('ready', callback);
    } else {
      console.error('onPlayerReady is not applied to player. Player is not initialized.');
    }
  }

  async onPlayerStateChanged(callback: Spotify.PlaybackStateListener): Promise<void> {
    if (this.player !== undefined) {
      this.player.removeListener('player_state_changed');
      this.player.addListener('player_state_changed', callback);
    } else {
      console.error('onPlayerStateChanged is not applied to player. Player is not initialized');
    }
  }

  async connect(): Promise<void> {
    if (this.player !== undefined) {
      await this.player.connect();
    }
  }

  async resume(): Promise<void> {
    if (this.player !== undefined) {
      try {
        await this.player.resume();
      } catch (error) {
        console.error(error);
      }
    }
  }

  async pause(): Promise<void> {
    if (this.player !== undefined) {
      try {
        await this.player.pause();
      } catch (error) {
        console.error(error);
      }
    }
  }

  async previousTrack(): Promise<void> {
    if (this.player !== undefined) {
      try {
        await this.player.previousTrack();
      } catch (error) {
        console.log(error);
      }
    }
  }

  async nextTrack(): Promise<void> {
    if (this.player !== undefined) {
      try {
        await this.player.nextTrack();
      } catch (error) {
        console.log(error);
      }
    }
  }
}

function connectPlayer(token: string): SpotifyPlayer {
  if (window.Spotify) {
    let player = new window.Spotify.Player({
      name: 'Chlorine',
      getOAuthToken: (cb) => {
        cb(token);
      },
    });

    player.addListener('initialization_error', ({ message }) => {
      console.error(message);
    });
    player.addListener('authentication_error', ({ message }) => {
      console.error(message);
    });
    player.addListener('account_error', ({ message }) => {
      console.error(message);
    });
    player.addListener('playback_error', ({ message }) => {
      console.error(message);
    });

    // Playback status updates
    player.addListener('player_state_changed', (status) => console.log(status));

    // Ready
    player.addListener('ready', ({ device_id }) => {
      console.log('Ready with Device ID', device_id);
    });

    // Not Ready
    player.addListener('not_ready', ({ device_id }) => {
      console.log('Device ID has gone offline', device_id);
    });

    return new SpotifyPlayer(player);
  }
}

export { connectPlayer, SpotifyPlayer };
