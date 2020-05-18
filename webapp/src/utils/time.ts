import { format } from 'date-fns';

function toTrackTime(milliseconds) {
  if (isNaN(milliseconds)) {
    return '00:00';
  }
  return format(milliseconds, 'mm:ss');
}

export { toTrackTime };
