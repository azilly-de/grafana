import { OrgRole } from '@grafana/data';

export enum TeamPermissionLevel {
  Admin = 4,
  Member = 0,
}

export { OrgRole as OrgRole };

export type PermissionLevel = 'view' | 'edit' | 'admin';

/** @deprecated Use PermissionLevel instead */
export enum PermissionLevelString {
  View = 'View',
  Edit = 'Edit',
  Admin = 'Admin',
}

export enum SearchQueryType {
  Folder = 'dash-folder',
  Dashboard = 'dash-db',
}
