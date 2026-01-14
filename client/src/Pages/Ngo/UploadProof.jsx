import { useEffect, useState } from "react";
import QRCode from "react-qr-code";

import { API_ENDPOINTS, WS_BASE_URL } from '@/config/api';
import { ENV } from '@/config/environment';

export default function UploadProof() {
  const [session, setSession] = useState(null);
  const [captures, setCaptures] = useState([]);
  const [error, setError] = useState(null);
  const [loading, setLoading] = useState(true);

  // Create proof session
  useEffect(() => {
    const createSession = async () => {
      try {
        setLoading(true);
        setError(null);
        const response = await fetch(API_ENDPOINTS.CREATE_PROOF_SESSION, {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
        });

        if (!response.ok) {
          const errorText = await response.text();
          throw new Error(`Failed to create session: ${errorText}`);
        }

        const data = await response.json();
        setSession(data.sessionId);
      } catch (err) {
        console.error("Error creating session:", err);
        setError(err.message);
      } finally {
        setLoading(false);
      }
    };

    createSession();
  }, []);

  // WebSocket listener
  useEffect(() => {
    if (!session) return;

    const ws = new WebSocket(
      `${WS_BASE_URL}/ws/proof/${session}`
    );

    ws.onopen = () => {
      console.log("WebSocket connected");
    };

    ws.onmessage = (event) => {
      try {
        const data = JSON.parse(event.data);
        console.log("Received WebSocket data:", data);
        console.log("Location data - lat:", data.lat, "lng:", data.lng);
        setCaptures((prev) => [...prev, data]);
      } catch (err) {
        console.error("Error parsing WebSocket message:", err);
      }
    };

    ws.onerror = (error) => {
      console.error("WebSocket error:", error);
    };

    ws.onclose = () => {
      console.log("WebSocket disconnected");
    };

    return () => {
      if (ws.readyState === WebSocket.OPEN) {
        ws.close();
      }
    };
  }, [session]);

  if (loading) {
    return (
      <div className="p-6">
        <p>Creating session...</p>
      </div>
    );
  }

  if (error) {
    return (
      <div className="p-6">
        <div className="text-red-600">Error: {error}</div>
        <button
          onClick={() => window.location.reload()}
          className="mt-4 px-4 py-2 bg-blue-500 text-white rounded"
        >
          Retry
        </button>
      </div>
    );
  }

  if (!session) {
    return (
      <div className="p-6">
        <p>No session available</p>
      </div>
    );
  }

  const qrUrl = `${ENV.FRONTEND_BASE_URL}/mobile/proof/${session}`;
  const apiBaseUrl = ENV.API_BASE_URL;

  return (
    <div className="p-6 max-w-4xl mx-auto">
      <h1 className="text-2xl font-bold mb-4">Upload Proof</h1>
      <p className="text-gray-600 mb-6">
        Scan this QR code with your mobile device to capture proof images with location and timestamp.
      </p>

      <div className="mt-4 bg-white p-4 inline-block border-2 border-gray-300 rounded-lg">
        <QRCode value={qrUrl} size={256} />
      </div>

      <div className="mt-4">
        <p className="text-sm text-gray-500">
          Session ID: <code className="bg-gray-100 px-2 py-1 rounded">{session}</code>
        </p>
        <p className="text-sm text-gray-500 mt-2">
          URL: <code className="bg-gray-100 px-2 py-1 rounded">{qrUrl}</code>
        </p>
      </div>

      <div className="mt-8">
        <h2 className="text-xl font-semibold mb-4">
          Captured Images ({captures.length})
        </h2>
        {captures.length === 0 ? (
          <p className="text-gray-500">No images captured yet. Scan the QR code to start capturing.</p>
        ) : (
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
            {captures.map((c, i) => (
              <div key={i} className="border rounded-lg p-3 shadow-sm">
                <img
                  src={`${apiBaseUrl}/uploads/${c.image}`}
                  alt={`Proof ${i + 1}`}
                  className="w-full h-48 object-cover rounded"
                  onError={(e) => {
                    e.target.src = "data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' width='200' height='200'%3E%3Crect fill='%23ddd' width='200' height='200'/%3E%3Ctext fill='%23999' font-family='sans-serif' font-size='14' x='50%25' y='50%25' text-anchor='middle' dy='.3em'%3EImage not found%3C/text%3E%3C/svg%3E";
                  }}
                />
                <div className="mt-2 text-sm">
                  <p className="font-medium">Location:</p>
                  <p className="text-gray-600">
                    {c.lat && c.lat !== "" && c.lng && c.lng !== "" ? (
                      <>
                        Lat: {parseFloat(c.lat).toFixed(6)}, Lng: {parseFloat(c.lng).toFixed(6)}
                      </>
                    ) : (
                      "N/A (Location not available)"
                    )}
                  </p>
                  <p className="text-gray-500 mt-1">
                    {c.timestamp ? new Date(c.timestamp).toLocaleString() : "No timestamp"}
                  </p>
                </div>
              </div>
            ))}
          </div>
        )}
      </div>
    </div>
  );
}
