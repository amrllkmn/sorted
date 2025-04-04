export namespace tasks {
	export class Task {
		ID: number;
		// Go type: time
		CreatedAt: any;
		// Go type: time
		UpdatedAt: any;
		// Go type: gorm
		DeletedAt: any;
		id: number;
		description: string;
		urgent?: boolean;
		important?: boolean;

		static createFrom(source: any = {}) {
			return new Task(source);
		}

		constructor(source: any = {}) {
			if ('string' === typeof source) source = JSON.parse(source);
			this.ID = source['ID'];
			this.CreatedAt = this.convertValues(source['CreatedAt'], null);
			this.UpdatedAt = this.convertValues(source['UpdatedAt'], null);
			this.DeletedAt = this.convertValues(source['DeletedAt'], null);
			this.id = source['id'];
			this.description = source['description'];
			this.urgent = source['urgent'];
			this.important = source['important'];
		}

		convertValues(a: any, classs: any, asMap: boolean = false): any {
			if (!a) {
				return a;
			}
			if (a.slice && a.map) {
				return (a as any[]).map((elem) => this.convertValues(elem, classs));
			} else if ('object' === typeof a) {
				if (asMap) {
					for (const key of Object.keys(a)) {
						a[key] = new classs(a[key]);
					}
					return a;
				}
				return new classs(a);
			}
			return a;
		}
	}
}
