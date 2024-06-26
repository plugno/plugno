import { WEBSOCKET_URL } from "@utils/env";
import React, { createContext, useContext, useEffect, useState } from "react";
import { useDispatch } from "react-redux";
import { Dispatch } from "../store";

export enum TrackingStatus {
  Accepted = "accepted",
  InTransit = "in_transit",
  Active = "active",
  Completed = "completed",
}

type SendSocketMessage = {
  type: TrackingStatus;
  jobId: number;
};

type TrackingContextProps = {
  sendTrackingMessage: (socketMessage: SendSocketMessage) => void;
};

const TrackingContext = createContext<TrackingContextProps | null>(null);

type TrackingSocketResponse = {
  jobId: number;
  type: TrackingStatus;
};

export const TrackingProvider = ({
  children,
}: {
  children: React.ReactNode;
}) => {
  const [trackingSocket, setTrackingSocket] = useState<WebSocket | null>(null);
  const dispatch = useDispatch<Dispatch>();

  // const handleOpenSocket = () => {};

  useEffect(() => {
    const _socket = new WebSocket(`ws://${WEBSOCKET_URL}/jobs/tracking`);

    setTrackingSocket(_socket);
  }, []);

  useEffect(() => {
    if (trackingSocket) {
      trackingSocket.onopen = () => {
        console.log("Tracking socket open");
      };

      trackingSocket.onmessage = (msg) => {
        const msgData = JSON.parse(msg.data) as TrackingSocketResponse;

        dispatch.jobs.updateTrackingStatus(msgData.type);
      };

      trackingSocket.onclose = () => {
        console.log("Tracking socket is closed");
      };
    }

    return () => {
      trackingSocket?.close(1000, "Closed chat page");
      console.log("TrackingProvider unmounted");
    };
  }, [trackingSocket, dispatch.jobs]);

  const handleSendSocketMessage = (socketMessage: SendSocketMessage) => {
    trackingSocket?.send(JSON.stringify(socketMessage));
  };

  /* useEffect(() => {
    fetch("http://localhost:6001/jobs/findActive", {
      method: "GET",
    })
      .then((res) => res.json())
      .then((data) => {
        console.log(data);

        handleOpenSocket();
      });
  }, []); */

  return (
    <TrackingContext.Provider
      value={{ sendTrackingMessage: handleSendSocketMessage }}
    >
      {children}
    </TrackingContext.Provider>
  );
};

export const useTrackingContext = () =>
  useContext<TrackingContextProps>(TrackingContext);
