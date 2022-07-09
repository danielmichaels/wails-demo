export namespace main {
	
	export class User {
	    avatar_url: string;
	    bio: string;
	    blog: string;
	    company: string;
	    // Go type: time.Time
	    created_at: any;
	    email: any;
	    events_url: string;
	    followers: number;
	    followers_url: string;
	    following: number;
	    following_url: string;
	    gists_url: string;
	    gravatar_id: string;
	    hireable: boolean;
	    html_url: string;
	    id: number;
	    location: string;
	    login: string;
	    name: string;
	    node_id: string;
	    organizations_url: string;
	    public_gists: number;
	    public_repos: number;
	    received_events_url: string;
	    repos_url: string;
	    site_admin: boolean;
	    starred_url: string;
	    subscriptions_url: string;
	    twitter_username: any;
	    type: string;
	    // Go type: time.Time
	    updated_at: any;
	    url: string;
	
	    static createFrom(source: any = {}) {
	        return new User(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.avatar_url = source["avatar_url"];
	        this.bio = source["bio"];
	        this.blog = source["blog"];
	        this.company = source["company"];
	        this.created_at = this.convertValues(source["created_at"], null);
	        this.email = source["email"];
	        this.events_url = source["events_url"];
	        this.followers = source["followers"];
	        this.followers_url = source["followers_url"];
	        this.following = source["following"];
	        this.following_url = source["following_url"];
	        this.gists_url = source["gists_url"];
	        this.gravatar_id = source["gravatar_id"];
	        this.hireable = source["hireable"];
	        this.html_url = source["html_url"];
	        this.id = source["id"];
	        this.location = source["location"];
	        this.login = source["login"];
	        this.name = source["name"];
	        this.node_id = source["node_id"];
	        this.organizations_url = source["organizations_url"];
	        this.public_gists = source["public_gists"];
	        this.public_repos = source["public_repos"];
	        this.received_events_url = source["received_events_url"];
	        this.repos_url = source["repos_url"];
	        this.site_admin = source["site_admin"];
	        this.starred_url = source["starred_url"];
	        this.subscriptions_url = source["subscriptions_url"];
	        this.twitter_username = source["twitter_username"];
	        this.type = source["type"];
	        this.updated_at = this.convertValues(source["updated_at"], null);
	        this.url = source["url"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
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

