import * as React from 'react';
import styled from 'styled-components';
import { useTranslation } from 'react-i18next';
import Button from './common/Button';
import TextSpan from './common/TextSpan';
import TextInput from './common/TextInput';
import { HorizontalFlex } from '../containers/common/HorizontalFlex';
import { useState } from 'react';
import Modal from './common/Modal';
import LinkButton from './common/LinkButton';
import RoomLinkModal from '../containers/RoomLinkModal';

interface SettingsProps {
  roomId: number;
}

const Settings: React.FC<SettingsProps> = ({ roomId }) => {
  const { t } = useTranslation();
  const [songsPerMember, setSongsPerMember] = useState<number>(0);
  const [maxMembers, setMaxMembers] = useState<number>(0);
  const [showLinkModal, setLinkModal] = useState<boolean>(false);

  return (
    <SettingsContainer>
      <TextSpan>
        {t('your_room')} {roomId ? roomId : ''}
      </TextSpan>
      <ValueSettingsContainer>
        <HorizontalFlex justifyContent='space-between' width='15rem'>
          <TextSpan>{t('songs_per_member')}</TextSpan>
          <TextInput
            onChange={(event) => {
              const targetValue = Number(event.currentTarget.value);
              if (!isNaN(targetValue)) {
                setSongsPerMember(targetValue);
              }
            }}
            width='2rem'
            value={songsPerMember.toString()}
          />
        </HorizontalFlex>
        <HorizontalFlex justifyContent='space-between' width='15rem'>
          <TextSpan>{t('max_members')}</TextSpan>
          <TextInput
            onChange={(event) => {
              const targetValue = Number(event.currentTarget.value);
              if (!isNaN(targetValue)) {
                setMaxMembers(targetValue);
              }
            }}
            width='2rem'
            value={maxMembers.toString()}
          />
        </HorizontalFlex>
      </ValueSettingsContainer>
      <div>
        <GetRoomLinkButton
          onClick={() => {
            setLinkModal(true);
          }}
        >
          {t('link')}
        </GetRoomLinkButton>
      </div>
      <RoomLinkModal visibility={showLinkModal} close={setLinkModal} roomId={roomId} />
    </SettingsContainer>
  );
};

const SettingsContainer = styled.div`
  display: flex;
  flex-direction: column;
  color: white;
  height: 100%;
  background-color: #292929;
  border-left: 1px dashed #616467;
  width: 20rem;
  align-items: center;

  & > * {
    padding-top: 1.5rem;
  }
`;

const ValueSettingsContainer = styled.div`
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  align-items: baseline;
`;

const GetRoomLinkButton = styled(Button)`
  font-size: 1rem;
  padding: 8px 24px;
`;

export default Settings;
