export type TTask = {
	id: number;
	description: string;
	urgent: boolean | null;
	important: boolean | null;
}

export const tasks : TTask[] = [
	{ id: 1, description: 'item1', urgent: null, important: null},
	{ id: 2, description: 'item2', urgent: null, important: null},
	{ id: 3, description: 'item3', urgent: null, important: null},
	{ id: 4, description: 'item4', urgent: null, important: null},
	{ id: 5, description: 'item5', urgent: null, important: null},
	{ id: 6, description: 'item6', urgent: null, important: null},
	{ id: 7, description: 'item7', urgent: null, important: null},
	{ id: 8, description: 'item8', urgent: null, important: null},
	{ id: 9, description: 'item9', urgent: null, important: null},
	{ id: 10, description: 'item10', urgent: null, important: null},
	{ id: 11, description: 'item11', urgent: false, important: true},
	{ id: 12, description: 'item12', urgent: false, important: true},
	{ id: 13, description: 'item13', urgent: true, important: false},
	{ id: 14, description: 'item14', urgent: true, important: false},
	{ id: 15, description: 'item15', urgent: false, important: false},
	{ id: 16, description: 'item16', urgent: false, important: false}
];
