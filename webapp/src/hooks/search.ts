import { useEffect, useState } from 'react';
import { ChlorineService } from '../services/chlorineService';
import { SpotifyTrack } from '../models/chlorine';

function useSongSearch() {
  const [songQuery, setSongQuery] = useState<string>('');
  const [searchResult, setSearchResult] = useState<SpotifyTrack[]>([]);

  useEffect(() => {
    async function prepare() {
      if (songQuery) {
        try {
          const songs = await new ChlorineService().searchTracks(songQuery);
          setSearchResult(songs);
        } catch (error) {
          console.log('Song search error');
          console.error(error);
        }
      }
    }

    prepare();
  }, [songQuery]);

  return { searchResult, setSongQuery };
}

export { useSongSearch };
