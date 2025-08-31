import Home from "../pages/Home";
import Quiz from "../pages/Quiz";
import Test from "../pages/Test";

export interface RouteConfig{
    to: string,
    primary: string
    element: React.ElementType
}

export const routes: RouteConfig[] = [
    {to: '/', primary: 'Dashboard', element: Home},
    {to: '/test', primary: 'Test', element: Test},
    {to: '/quiz', primary: 'Quiz', element: Quiz}

]