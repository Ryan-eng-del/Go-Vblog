package blog

type Blog struct {
	// 文章Id
	Id int
	// 创建时间
	CreateAt int64
	// 更新时间
	UpdateAt int64
	// 发布时间
	PublishAt int64
	// 用户提交数据
	*CreateBlogRequest
	// 文章状态 草稿/发布
	Status Status
}

type BlogSet struct {
	blogs []*Blog
	total int32
}

type CreateBlogRequest struct {
	// 文章摘要信息,通过提前Content内容获取
	Summary string
	// 文章图片
	TitleImg string
	// 文章标题
	TitleName string
	// 文章副标题
	SubTitle string
	// 文章内容
	Content string
	// 文章作者
	Author string
}

type UpdateBlogRequest struct {
	BlogId     int
	UpdateMode UpdateMode
	*CreateBlogRequest
}

type DeleteBlogRequest struct {
	Id int
}

type QueryBlogRequest struct {
	PageSize   int
	PageNumber int
	Keywords   string
	// 补充状态过滤参数, 用于web 前台 过滤已经发布的文章
	// 比如过滤 状态为发布的文章
	Status  *Status
	BlogIds []int
}

type DescribeBlogRequest struct {
	Id int
}

type UpdateBlogStatusRequest struct {
	// 文章Id
	Id int
	// 文章状态 草稿/发布
	Status Status
}
