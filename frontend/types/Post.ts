export type Post = {
	id: number
	post_title: string
	post_content: string
	post_excerpt?: string
	post_date: string
	post_author: string
	post_status: string
	comment_status: string
	comment_count: number
	post_name: string
	post_modified: string
	post_parent: number
	guid: string
	post_type: 'post' | 'page'
	group_comments: PostComment[]
}

export type PostComment = {
	comment_id: number;
	comment_post_id: number;
	comment_author: string;
	comment_author_email: string;
	comment_author_url: string;
	comment_author_ip: string;
	comment_date: string;
	comment_date_gmt: string;
	comment_content: string;
	comment_karma: number;
	comment_agent: string;
	comment_type: string;
	comment_parent: number;
	user_id: number;
	children?: PostComment[] | null;
}