import { useState, useRef, useEffect } from "react";
import { useParams } from "react-router-dom";
import { API_ENDPOINTS } from '@/config/api';

export default function MobileProofCapture() {
  const { sessionID } = useParams();
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(null);
  const [success, setSuccess] = useState(false);
  const [location, setLocation] = useState(null);
  const fileInputRef = useRef(null);

  // Get location when component mounts
  useEffect(() => {
    if (navigator.geolocation) {
      navigator.geolocation.getCurrentPosition(
        (position) => {
          console.log("Initial location captured:", position.coords.latitude, position.coords.longitude);
          setLocation({
            lat: position.coords.latitude,
            lng: position.coords.longitude,
          });
        },
        (err) => {
          console.warn("Geolocation error on mount:", err.code, err.message);
          // Error codes: 1=PERMISSION_DENIED, 2=POSITION_UNAVAILABLE, 3=TIMEOUT
          if (err.code === 1) {
            console.warn("Location permission denied by user");
          } else if (err.code === 2) {
            console.warn("Location unavailable");
          } else if (err.code === 3) {
            console.warn("Location request timed out");
          }
        },
        { 
          timeout: 15000, 
          maximumAge: 0,
          enableHighAccuracy: true 
        }
      );
    } else {
      console.warn("Geolocation API not available in this browser");
    }
  }, []);

  const handleFileSelect = async (event) => {
    const file = event.target.files?.[0];
    if (!file) return;

    try {
      setLoading(true);
      setError(null);
      setSuccess(false);

      // Get current location (refresh if needed)
      let position = null;
      if (navigator.geolocation) {
        try {
          position = await new Promise((res, rej) => {
            navigator.geolocation.getCurrentPosition(
              res,
              rej,
              { 
                timeout: 10000, 
                maximumAge: 0,
                enableHighAccuracy: true 
              }
            );
          });
          console.log("Location captured:", position.coords.latitude, position.coords.longitude);
          setLocation({
            lat: position.coords.latitude,
            lng: position.coords.longitude,
          });
        } catch (geoError) {
          console.warn("Geolocation error:", geoError);
          // Use previously stored location if available
          if (location && location.lat && location.lng) {
            console.log("Using cached location:", location);
            position = {
              coords: {
                latitude: location.lat,
                longitude: location.lng,
              },
            };
          }
        }
      }

      // Create form data (timestamp = capture/upload time for server validation)
      const form = new FormData();
      const now = new Date();
      form.append("file", file, `proof-${now.getTime()}.jpg`);
      form.append("timestamp", now.toISOString());
      
      let latValue = "";
      let lngValue = "";
      
      if (position && position.coords) {
        latValue = position.coords.latitude.toString();
        lngValue = position.coords.longitude.toString();
        console.log("Sending location in form:", latValue, lngValue);
      } else if (location && location.lat && location.lng) {
        latValue = location.lat.toString();
        lngValue = location.lng.toString();
        console.log("Sending cached location in form:", latValue, lngValue);
      } else {
        console.warn("No location available to send");
      }
      
      form.append("lat", latValue);
      form.append("lng", lngValue);

      // Upload image
      const response = await fetch(
        API_ENDPOINTS.UPLOAD_PROOF_IMAGE(sessionID),
        {
          method: "POST",
          body: form,
        }
      );

      if (!response.ok) {
        const errorText = await response.text();
        throw new Error(`Upload failed: ${errorText}`);
      }

      setSuccess(true);
      setTimeout(() => {
        setSuccess(false);
      }, 3000);

      // Reset file input for next capture
      if (fileInputRef.current) {
        fileInputRef.current.value = "";
      }
    } catch (err) {
      console.error("Error uploading:", err);
      setError(err.message || "Failed to upload image");
    } finally {
      setLoading(false);
    }
  };

  const triggerFileInput = () => {
    if (fileInputRef.current) {
      fileInputRef.current.click();
    }
  };

  return (
    <div className="min-h-screen bg-gradient-to-br from-blue-50 to-indigo-100 p-6 flex flex-col items-center justify-center">
      <div className="bg-white rounded-lg shadow-lg p-8 max-w-md w-full">
        <h1 className="text-2xl font-bold text-center mb-2">
          Capture Proof Image
        </h1>
        <p className="text-gray-600 text-center mb-6 text-sm">
          Session: <code className="bg-gray-100 px-2 py-1 rounded text-xs">{sessionID}</code>
        </p>

        {location && location.lat && location.lng ? (
          <div className="mb-4 p-3 bg-green-50 border border-green-200 rounded-lg">
            <p className="text-sm text-gray-700">
              <strong>✓ Location Ready:</strong> {location.lat.toFixed(6)}, {location.lng.toFixed(6)}
            </p>
          </div>
        ) : (
          <div className="mb-4 p-3 bg-yellow-50 border border-yellow-200 rounded-lg">
            <p className="text-sm text-yellow-700">
              <strong>⚠ Location:</strong> Not available. Please allow location access in your browser settings.
            </p>
          </div>
        )}

        {error && (
          <div className="mb-4 p-3 bg-red-50 border border-red-200 rounded-lg">
            <p className="text-sm text-red-700">{error}</p>
          </div>
        )}

        {success && (
          <div className="mb-4 p-3 bg-green-50 border border-green-200 rounded-lg">
            <p className="text-sm text-green-700">
              ✓ Proof uploaded successfully!
            </p>
          </div>
        )}

        {/* Hidden file input */}
        <input
          ref={fileInputRef}
          type="file"
          accept="image/*"
          capture="environment"
          onChange={handleFileSelect}
          className="hidden"
        />

        <button
          onClick={triggerFileInput}
          disabled={loading}
          className={`w-full py-4 px-6 rounded-lg font-semibold text-white transition-all ${
            loading
              ? "bg-gray-400 cursor-not-allowed"
              : "bg-green-600 hover:bg-green-700 active:bg-green-800"
          }`}
        >
          {loading ? (
            <span className="flex items-center justify-center">
              <svg
                className="animate-spin -ml-1 mr-3 h-5 w-5 text-white"
                xmlns="http://www.w3.org/2000/svg"
                fill="none"
                viewBox="0 0 24 24"
              >
                <circle
                  className="opacity-25"
                  cx="12"
                  cy="12"
                  r="10"
                  stroke="currentColor"
                  strokeWidth="4"
                ></circle>
                <path
                  className="opacity-75"
                  fill="currentColor"
                  d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"
                ></path>
              </svg>
              Uploading...
            </span>
          ) : (
            "Capture & Upload"
          )}
        </button>

        <p className="text-xs text-gray-500 text-center mt-4">
          Tap the button to open your camera and capture a photo. Location will be automatically included.
        </p>
      </div>
    </div>
  );
}
