import TextSpan from '../components/common/TextSpan';
import LinkButton from '../components/common/LinkButton';
import Modal from '../components/common/Modal';
import * as React from 'react';
import { useTranslation } from 'react-i18next';
import { AlignedCenterVerticalFlex } from './common/AlignedCenterFlex';
import styled from 'styled-components';

interface RoomLinkModalProps {
  visibility: boolean;
  close(newValue: boolean): void;
  roomId: number;
}

const RoomLinkModal: React.FC<RoomLinkModalProps> = ({ visibility, close, roomId }) => {
  const { t } = useTranslation();
  return (
    <Modal title={t('invitation_link')} display={[visibility, close]}>
      <AlignedCenterVerticalFlex>
        <TextSpan>{t('invitation_modal_description')}</TextSpan>
        <RoomLink
          onClick={() => {
            navigator.clipboard.writeText(
              `${window.location.protocol}//${window.location.host}/join?room=${roomId}`
            );
          }}
        >{`${window.location.protocol}//${window.location.host}/join?room=${roomId}`}</RoomLink>
      </AlignedCenterVerticalFlex>
    </Modal>
  );
};

const RoomLink = styled(LinkButton)`
  margin-top: 1rem;
  padding: 8px;
  border: 1px dashed #616467;
`;

export default RoomLinkModal;
