import { AppShell } from "@mantine/core";
import { useContext, useMemo } from "react";
import FiltersDrawer from "../components/FiltersDrawer";
import { AppContext } from "../context/AppContext";
import Menu from "../components/Menu";
import BasicMap from "./BasicMap";
import OffersListView from "./OffersListView";
import { useQueries, useQuery } from "@tanstack/react-query";
import AdminConsole from "./AdminConsole";
import { useGeolocation } from "./../hooks/useGeolocation";
import UserProfile from "./UserProfile/UserProfile";
import { BASE_URL } from "../shared/api/config";

export default function Home() {
  const { viewType } = useContext(AppContext);

  const { latitude, longitude } = useGeolocation();

  const { data } = useQuery({
    queryKey: ["offers"],
    queryFn: () =>
      fetch(`${BASE_URL}/announcements/filter`, {
        method: "POST", // *GET, POST, PUT, DELETE, etc.
        // mode: "cors", // no-cors, *cors, same-origin
        // cache: "no-cache", // *default, no-cache, reload, force-cache, only-if-cached
        // credentials: "same-origin", // include, *same-origin, omit
        headers: {
          "Content-Type": "application/json",
          // 'Content-Type': 'application/x-www-form-urlencoded',
        },
        // redirect: "follow", // manual, *follow, error
        // referrerPolicy: "no-referrer", // no-referrer, *no-referrer-when-downgrade, origin, origin-when-cross-origin, same-origin, strict-origin, strict-origin-when-cross-origin, unsafe-url
        body: JSON.stringify({ latitude, longitude, sort_by: "distance" }), // body data type must match "Content-Type" header
      }).then((res) => res.json()),
  });

  console.log("DATA");
  console.log(data);

  const queries = useQueries({
    queries: data?.announcements
      ? data.announcements.map((offer) => {
          return {
            queryKey: ["companies", offer.company_id],
            queryFn: () =>
              fetch(
                `http://localhost:8888/backend/companies/${offer.company_id}`
              ).then((res) => res.json()),
          };
        })
      : [], // if data?.announcements is undefined, an empty array will be returned
    combine: (results) => {
      const companies = results.reduce((acc, item) => {
        acc[item.data?.company_id] = item.data;
        return acc;
      }, {});

      return data?.announcements?.map((offer) => {
        return { ...offer, companyData: companies[offer.company_id] };
      });
      // res: companies
    },
  });

  const smth = useMemo(() => {
    console.log("MEMO!");
    console.log(queries);

    return queries;
  }, [queries]);

  return (
    <AppShell header={{ height: 80, offset: true }}>
      <AppShell.Header style={{ alignContent: "center" }}>
        <Menu></Menu>
      </AppShell.Header>

      <AppShell.Main>
        {viewType === "map" && (
          <>
            <BasicMap offers={smth}></BasicMap>
            <FiltersDrawer></FiltersDrawer>
          </>
        )}
        {viewType === "list" && (
          <>
            <OffersListView offers={smth}></OffersListView>
          </>
        )}
        {viewType === "admin" && <AdminConsole></AdminConsole>}
        {viewType === "profile" && <UserProfile />}
  

      </AppShell.Main>
    </AppShell>
  );
}
