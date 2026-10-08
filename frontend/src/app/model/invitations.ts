export interface Invitations {
  groupId: string;
  leaderName: string;
  members: string[];
  status?: 'PENDING' | 'ACCEPTED' | 'REJECTED';
}