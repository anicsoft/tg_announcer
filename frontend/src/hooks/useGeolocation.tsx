import {useEffect, useState} from 'react'


export const useGeolocation = () => {
  const [latitude, setLatitude] = useState<number>(0);
  const [longitude, setLongitude] = useState<number>(0);
  const [locationInfo, setLocationInfo] = useState({})
  const [locationError, setLocationError] = useState()

  useEffect(() => {
    const { geolocation } = navigator;
    
    const onSuccess = (res: any) => {
      setLatitude(res.coords.latitude);
      setLongitude(res.coords.longitude);
    }

    const onError = (res: any) => {
      setLocationError(res.message)
    }

    const watcherId = navigator.geolocation.watchPosition(onSuccess, onError);

    return () => {
      navigator.geolocation.clearWatch(watcherId);
    };
  
    // if (!locationError && !locationInfo) {
    //   if (navigator.permissions) {
    //     // Ask permissions in browser
    //     navigator.permissions.query({ name: "geolocation" }).then((result) => {
    //       if (result.state === "prompt") {
    //         geolocation.getCurrentPosition(onSuccess, onError);
    //       }
    //     })
    //   } else {
    //     // For Telegram bot. Permissions are asked by telegram itself
    //     geolocation.getCurrentPosition(onSuccess, onError);
    //   }
    // }
  }, [ ])

  return {latitude, longitude,locationError}
}
