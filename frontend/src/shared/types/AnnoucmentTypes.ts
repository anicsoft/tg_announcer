interface Announcement {
    id: number;
    company_id: number;
  }
  
  export interface CardProps {
    id: number;
  }
  
 export  interface AnnouncementsResponse {
    announcements: Announcement[];
  }
  