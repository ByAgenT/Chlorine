import * as React from 'react';
import TextInput from '../components/common/TextInput';
import SongSearchResultList from '../components/SongSearchResultList';
import Modal from '../components/common/Modal';
import { useTranslation } from 'react-i18next';
import { useSongSearch } from '../hooks/search';
import debounce from 'lodash/debounce';
import { useCallback } from 'react';

interface AddSongsModalProps {
  isShowed: boolean;
  onClose(value: boolean): void;
  onSongAdd(id: string): void;
}

const AddSongsModal: React.FC<AddSongsModalProps> = ({ isShowed, onClose, onSongAdd }) => {
  const { t } = useTranslation();
  const { searchResult, setSongQuery } = useSongSearch();

  const updateSongQuery = debounce((event: React.ChangeEvent<HTMLTextAreaElement>) => {
    setSongQuery(event.target.value);
  }, 200);

  const onSearchValueChange = useCallback(
    (event: React.ChangeEvent<HTMLTextAreaElement>) => {
      event.persist();
      updateSongQuery(event);
    },
    [updateSongQuery]
  );

  return (
    <Modal title={t('modal_title')} display={[isShowed, onClose]}>
      <TextInput placeholder={t('modal_search_placeholder')} onChange={onSearchValueChange} />
      <SongSearchResultList onSongAdd={onSongAdd} songs={searchResult} />
    </Modal>
  );
};

export default AddSongsModal;
