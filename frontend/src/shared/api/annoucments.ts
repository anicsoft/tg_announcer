import { useMutation, useQuery } from "@tanstack/react-query";
import { useContext } from "react";
import { AppContext } from "../../context/AppContext";
import { BASE_URL } from "../../shared/api/config";
import { useGeolocation } from "../../hooks/useGeolocation";

export const useFetchOffers = () => {
  const { initDataRaw } = useContext(AppContext);
  const { latitude, longitude } = useGeolocation();
console.log("initData",initDataRaw)
  const { data, refetch } = useQuery({
    queryKey: ["offers"],
    queryFn: () =>
      fetch(`${BASE_URL}/announcements/filter`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Authorization: `tma ${initDataRaw}`,
        },
        body: JSON.stringify({ latitude, longitude, sort_by: "distance" }),
      }).then((res) => res.json()),
  });

  return { data, refetch };
};


