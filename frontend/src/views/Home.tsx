import { AppShell } from "@mantine/core";
import { useContext, useMemo } from "react";
import FiltersDrawer from "../components/FiltersDrawer";
import { AppContext } from "../context/AppContext";
import Menu from "../components/Menu";
import BasicMap from "./BasicMap";
import OffersListView from "./OffersListView";
import { useQueries } from "@tanstack/react-query";
import AdminConsole from "./AdminConsole";
import UserProfile from "./UserProfile/UserProfile";
import { useFetchOffers } from "../shared/api/annoucments";

export default function Home() {
  const { viewType } = useContext(AppContext);


  const { data, refetch } = useFetchOffers();



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
