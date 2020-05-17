import * as React from 'react';
import TextInput from '../components/common/TextInput';
import SongSearchResultList from '../components/SongSearchResultList';
import Modal from '../components/common/Modal';
import { useTranslation } from 'react-i18next';
import { SpotifyTrack } from '../models/chlorine';

interface AddSongsModalProps {
  isShowed: boolean;
  onClose(value: boolean): void;
  onSearchValueChange(event: React.ChangeEvent<HTMLTextAreaElement>): void;
  onSongAdd(id: string): void;
  songs: SpotifyTrack[];
}

const AddSongsModal: React.FC<AddSongsModalProps> = ({
  isShowed,
  onClose,
  onSearchValueChange,
  onSongAdd,
  songs,
}) => {
  const { t } = useTranslation();
  return (
    <Modal title={t('modal_title')} display={[isShowed, onClose]}>
      <TextInput placeholder={t('modal_search_placeholder')} onChange={onSearchValueChange} />
      <SongSearchResultList onSongAdd={onSongAdd} songs={songs} />
    </Modal>
  );
};

export default AddSongsModal;
