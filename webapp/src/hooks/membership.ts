import { Member } from '../models/chlorine';
import { useEffect, useState } from 'react';
import { ChlorineService } from '../services/chlorineService';

function useMemberInformation(): [Member | null, () => void] {
  const [member, setMember] = useState<Member | null>(null);

  async function refreshMember() {
    try {
      setMember(await new ChlorineService().getMemberInfo());
    } catch (error) {
      console.error(error);
    }
  }

  useEffect(() => {
    async function prepare() {
      try {
        setMember(await new ChlorineService().getMemberInfo());
      } catch (error) {
        console.error(error);
      }
    }

    prepare();
  }, []);

  return [member, refreshMember];
}

function useMembersList(): [Member[], () => void] {
  const [members, setMembers] = useState<Member[]>([]);

  async function updateMembers() {
    try {
      const members = await new ChlorineService().getRoomMembers();
      setMembers(members);
    } catch (error) {
      console.error(error);
    }
  }

  useEffect(() => {
    async function prepare() {
      try {
        const members = await new ChlorineService().getRoomMembers();
        setMembers(members);
      } catch (error) {
        console.error(error);
      }
    }

    prepare();
  }, []);

  return [members, updateMembers];
}

export { useMemberInformation, useMembersList };
