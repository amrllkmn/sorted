// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
import { tasks } from '../models';

export function CreateTask(arg1: string): Promise<tasks.Task>;

export function DeleteTask(arg1: number): Promise<void>;

export function GetAllTasks(): Promise<Array<tasks.Task>>;

export function UpdateTaskStatus(arg1: number, arg2: boolean, arg3: boolean): Promise<void>;
