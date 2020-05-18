import { useEffect, useState } from 'react';

interface WSResponse {
  type: 'response' | 'broadcast';
  status: 'ok' | 'error';
  description: string;
  body: Record<string, Object>;
}

class ChlorineWebSocketConnection {
  private connection: WebSocket;
  private readonly broadcastCallbackMap: Map<string, (response: WSResponse) => void>;

  constructor() {
    this.broadcastCallbackMap = new Map<string, (response: WSResponse) => void>();
  }

  public establishOnce() {
    if (!this.connection) {
      this.connection = new WebSocket(`ws://${window.location.host}/ws`);
      this.connection.onerror = this.error.bind(this);
      this.connection.onmessage = this.listen.bind(this);
    }
  }

  private error(event: Event): void {
    console.error('WebSocket error');
    console.error(event);
  }

  private listen(event: MessageEvent): void {
    const response: WSResponse = JSON.parse(event.data);
    if (response.type === 'broadcast') {
      if (this.broadcastCallbackMap.has(response.description)) {
        this.broadcastCallbackMap.get(response.description)(response);
      }
    }
  }

  public onBroadcast(description: string, callback: (response: WSResponse) => void): void {
    this.broadcastCallbackMap.set(description, callback);
  }

  public removeOnBroadcastListener(description: string): void {
    this.broadcastCallbackMap.delete(description);
  }
}

const GlobalWebSocketConnection = new ChlorineWebSocketConnection();

function useChlorineWebSocket() {
  const [connection] = useState<ChlorineWebSocketConnection>(GlobalWebSocketConnection);

  useEffect(() => {
    connection.establishOnce();
  }, [connection]);

  return connection;
}

export default useChlorineWebSocket;
