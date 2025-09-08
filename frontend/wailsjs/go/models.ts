export namespace entity {
	
	export class Card {
	    ID: number;
	    DeckID: number;
	    Front: string;
	    Back: string;
	    // Go type: time
	    CreatedAt: any;
	    Repetitions: number;
	    Lapses: number;
	    EaseFactor: number;
	    Interval: number;
	    // Go type: time
	    Due: any;
	    // Go type: time
	    LastReviewed: any;
	
	    static createFrom(source: any = {}) {
	        return new Card(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ID = source["ID"];
	        this.DeckID = source["DeckID"];
	        this.Front = source["Front"];
	        this.Back = source["Back"];
	        this.CreatedAt = this.convertValues(source["CreatedAt"], null);
	        this.Repetitions = source["Repetitions"];
	        this.Lapses = source["Lapses"];
	        this.EaseFactor = source["EaseFactor"];
	        this.Interval = source["Interval"];
	        this.Due = this.convertValues(source["Due"], null);
	        this.LastReviewed = this.convertValues(source["LastReviewed"], null);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
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
	export class Deck {
	    ID: number;
	    Name: string;
	    Description: string;
	    // Go type: time
	    CreatedAt: any;
	    // Go type: time
	    UpdatedAt: any;
	    NewCardLimit: number;
	    ReviewLimit: number;
	    // Go type: time
	    LastStudiedAt: any;
	    Cards: Card[];
	
	    static createFrom(source: any = {}) {
	        return new Deck(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ID = source["ID"];
	        this.Name = source["Name"];
	        this.Description = source["Description"];
	        this.CreatedAt = this.convertValues(source["CreatedAt"], null);
	        this.UpdatedAt = this.convertValues(source["UpdatedAt"], null);
	        this.NewCardLimit = source["NewCardLimit"];
	        this.ReviewLimit = source["ReviewLimit"];
	        this.LastStudiedAt = this.convertValues(source["LastStudiedAt"], null);
	        this.Cards = this.convertValues(source["Cards"], Card);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
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

