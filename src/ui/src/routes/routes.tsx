import AddDisease from "../pages/AddDisease";
import DNATest from "../pages/DNATest";
import Results from "../pages/Results";

export type route = {
  label: string;
  path: string;
  component?: () => JSX.Element;
};

type navroutes = {
  content: route;
  children_routes?: route[];
  isDropdown?: boolean;
  parentPath?: string;
};

export const toRoute = (label: string, path: string, component?: (props?: any) => JSX.Element): route => ({
  label,
  path,
  component,
});

export const ADD_DISEASE_PAGE = toRoute("Add Disease", "/", AddDisease);
export const DNA_TEST_PAGE = toRoute("DNA Test", "/dna-test", DNATest);
export const RESULTS_PAGE = toRoute("results", "/results", Results);

export const AllRoutes = [ADD_DISEASE_PAGE, DNA_TEST_PAGE, RESULTS_PAGE];